import numpy as np
import matplotlib.pyplot as plt
from sklearn.cluster import DBSCAN
import matplotlib
from common import *


class DataGaner():

    def __init__(self, type="complex", size=250):
        self.size = size
        self.type = type
        self.create_data()

    def create_data(self):
        if self.type == "complex":
            self.data = [(np.random.randn()/6 - 1, np.random.randn()/6 - 1) for i in range(self.size)]
            self.data.extend([(np.random.randn()/4 + 2.5, np.random.randn()/5) for i in range(self.size)])
            self.data.extend([(np.random.randn()/3 - 2, np.random.randn()/10 + 1) for i in range(self.size)])
            self.data.extend([(np.random.randn()/50 - 2, np.random.randn() + 1) for i in range(self.siz