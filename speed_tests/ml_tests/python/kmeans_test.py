import numpy as np
from sklearn.cluster import KMeans
import matplotlib.pyplot as plt
from common import *


class KMeansTester:

    train = None,
    train_test = None,
    kmeans = None

    def __init__(self):
        pass

    def set_data(self, path: str, path_test: str):
        self.train = np.genfromtxt(path, delimiter=',')[1:]
        self.trai