import onnx
from google.protobuf.json_format import MessageToDict

model = onnx.load("example.onnx")
for _input in model.graph.input:
    dim = _input.type.tensor_type.shape.dim
    input_shape = [MessageToDict(d).get("dimValue") for d in dim]

print(input_shape)