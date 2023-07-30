# Specify Training Procedure for Docker Environment Here

# Run Python 3.8
FROM python:3.8

# Copy this entire folder into the training directory
COPY . /requestor
WORKDIR /requestor

RUN pip3 install --upgrade pip
RUN pip3 install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cpu
# Install dependencies for model
RUN pip3 install -r requirements.txt

# Convert ONNX File to PyTorch File 
RUN python3 convert.py 

# Run training script with data path
CMD python3 training.py
# Testing Purposes
# ENTRYPOINT ["tail", "-f", "/dev/null"]