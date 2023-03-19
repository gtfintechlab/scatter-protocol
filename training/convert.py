import torch
from onnx2torch import convert

# Path to ONNX model
onnx_model_path = './example.onnx'
torch_model = convert(onnx_model_path)
torch.save(torch_model, './example.pt')

model_loaded = torch.load('./example.pt')

print(model_loaded)