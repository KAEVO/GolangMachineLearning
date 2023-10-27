import numpy as np
from sklearn.decomposition import PCA
from common import *


class PCATester:
    X = None,
    pca = None

    def __init__(self):
        pass

    def set_data(self, path: str):
        self.X = np.genfromtxt(pat