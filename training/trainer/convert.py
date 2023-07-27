import torch
from onnx2torch import convert

# Path to ONNX model
onnx_model_path = './models/example.onnx'
torch_model = convert(onnx_model_path)
torch.save(torch_model, './models/example.pt')

model_loaded = torch.load('./models/example.pt')

print(model_loaded)