import grpc
import grpcpb.add_pb2 as add_pb2
import grpcpb.add_pb2_grpc as add_pb2_grpc


def validate_document(number, format):
    channel = grpc.insecure_channel('localhost:50051')
    stub = add_pb2_grpc.DocumentValidationStub(channel)
    request = add_pb2.DocumentRequest(number=number, format=format)
    response = stub.ValidateDocument(request)
    print(f"Received response: Number={response.number}, ResponseCode={response.response_code}")

if __name__ == '__main__':
    number = input("Digite o número do documento: ")
    format = input("Digite o formato (0 para CPF, 1 para CNPJ): ")
    validate_document(number, format)
