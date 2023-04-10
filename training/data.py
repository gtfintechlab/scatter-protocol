import glob
from pandas.core.common import flatten
import cv2
from torch.utils.data import Dataset
import albumentations as A
from albumentations.pytorch import ToTensorV2

augmentations = A.Compose(
    [
        A.SmallestMaxSize(max_size=700),
        A.ShiftScaleRotate(shift_limit=0.05, scale_limit=0.05, rotate_limit=360, p=0.5),
        A.RandomCrop(height=640, width=640),
        A.RGBShift(r_shift_limit=15, g_shift_limit=15, b_shift_limit=15, p=0.5),
        A.RandomBrightnessContrast(p=0.5),
        A.MultiplicativeNoise(multiplier=[0.5,2], per_channel=True, p=0.2),
        A.Normalize(mean=(0.485, 0.456, 0.406), std=(0.229, 0.224, 0.225)),
        A.HueSaturationValue(hue_shift_limit=0.2, sat_shift_limit=0.2, val_shift_limit=0.2, p=0.5),
        A.RandomBrightnessContrast(brightness_limit=(-0.1,0.1), contrast_limit=(-0.1, 0.1), p=0.5),
        ToTensorV2(),

    ]
)

class DataProcessor:
    def __init__(self, trainDataPath='') -> None:
        self.trainDataPath = trainDataPath
        self.trainImagePaths = []
        self.classes = []

        for data_path in glob.glob(self.trainDataPath + '/*'):
            self.classes.append(data_path.split('/')[-1]) 
            self.trainImagePaths.append(glob.glob(data_path + '/*'))

        self.trainImagePaths = list(flatten(self.trainImagePaths))

        self.indexToClass = {i:j for i, j in enumerate(self.classes)}
        self.classToIndex = {value:key for key,value in self.indexToClass.items()}
        

class CustomDataset(Dataset):
    def __init__(self, trainDataPath, transform=augmentations):
        self.dataloader = DataProcessor(trainDataPath)
        self.imagePaths = self.dataloader.trainImagePaths
        self.transform = transform

    def __len__(self):
        return len(self.imagePaths)
    
    def __getitem__(self, index):
        imagePath = self.imagePaths[index]
        image = cv2.imread(imagePath)
        image = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)
        
        label = imagePath.split('/')[-2]
        label = self.dataloader.classToIndex[label]

        if self.transform:
            image = self.transform(image=image)["image"]
        
        return image, label
