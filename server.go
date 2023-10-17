package main

import (
	"context"
	"log"
	"net"
	"strings"

	pb "GRPC/grpcpb"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) ValidateDocument(ctx context.Context, req *pb.DocumentRequest) (*pb.DocumentResponse, error) {
	number := req.Number
	format := req.Format

	if len(number) != 11 && len(number) != 14 {
		return &pb.DocumentResponse{Number: number, ResponseCode: 2}, nil // Tamanho de documento inválido
	}

	if !isValidFormat(number, format) {
		return &pb.DocumentResponse{Number: number, ResponseCode: 3}, nil // Formato inválido
	}

	if format == "0" && !isValidCPF(number) {
		return &pb.DocumentResponse{Number: number, ResponseCode: 1}, nil // CPF inválido
	} else if format == "1" && !isValidCNPJ(number) {
		return &pb.DocumentResponse{Number: number, ResponseCode: 1}, nil // CNPJ inválido
	}

	return &pb.DocumentResponse{Number: number, ResponseCode: 0}, nil // Válido
}

func isValidFormat(number, format string) bool {
	if (format == "0" && isValidCPF(number)) || (format == "1" && isValidCNPJ(number)) {
		return true
	}
	return false
}

func isValidCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	if len(cpf) != 11 {
		return false
	}

	// Verifique se todos os dígitos são iguais (situação inválida)
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			break
		}
		if i == 10 {
			return false
		}
	}

	// Valide o primeiro dígito verificador
	sum1 := 0
	for i := 0; i < 9; i++ {
		digit := int(cpf[i] - '0')
		sum1 += digit * (10 - i)
	}
	remainder1 := sum1 % 11
	if remainder1 < 2 {
		if int(cpf[9]-'0') != remainder1 {
			return false
		}
	} else {
		if int(cpf[9]-'0') != 11-remainder1 {
			return false
		}
	}

	// Valide o segundo dígito verificador
	sum2 := 0
	for i := 0; i < 10; i++ {
		digit := int(cpf[i] - '0')
		sum2 += digit * (11 - i)
	}
	remainder2 := sum2 % 11
	if remainder2 < 2 {
		if int(cpf[10]-'0') != remainder2 {
			return false
		}
	} else {
		if int(cpf[10]-'0') != 11-remainder2 {
			return false
		}
	}

	return true
}

func isValidCNPJ(cnpj string) bool {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	if len(cnpj) != 14 {
		return false
	}

	// Verifique se todos os dígitos são iguais (situação inválida)
	for i := 1; i < 14; i++ {
		if cnpj[i] != cnpj[0] {
			break
		}
		if i == 13 {
			return false
		}
	}

	// Valide o primeiro dígito verificador
	multipliers1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum1 := 0
	for i := 0; i < 12; i++ {
		digit := int(cnpj[i] - '0')
		sum1 += digit * multipliers1[i]
	}
	remainder1 := sum1 % 11
	if remainder1 < 2 {
		if int(cnpj[12]-'0') != remainder1 {
			return false
		}
	} else {
		if int(cnpj[12]-'0') != 11-remainder1 {
			return false
		}
	}

	// Valide o segundo dígito verificador
	multipliers2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum2 := 0
	for i := 0; i < 13; i++ {
		digit := int(cnpj[i] - '0')
		sum2 += digit * multipliers2[i]
	}
	remainder2 := sum2 % 11
	if remainder2 < 2 {
		if int(cnpj[13]-'0') != remainder2 {
			return false
		}
	} else {
		if int(cnpj[13]-'0') != 11-remainder2 {
			return false
		}
	}

	return true
}

func (s *server) mustEmbedUnimplementedDocumentValidationServer() {}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDocumentValidationServer(s, &server{})
	log.Println("Server is running on :50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
