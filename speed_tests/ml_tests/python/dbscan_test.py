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
            self.data.extend([(np.random.randn()/50 - 2, np.random.randn() + 1) for i in range(self.size)])
            self.data.extend([(np.random.randn()/5 + 1, np.random.randn()/2 + 1) for i in range(self.size)])
            self.data.extend([(i/25 - 1, + np.random.randn()/20 - 3) for i in range(self.size)])
            self.data.extend([(i/25 - 2.5, 9 - (i/50 - 2)**2 + np.random.randn()/20) for i in range(self.size)])
            self.data.extend([(i/25 - 2.5, 6 + (i/50 - 2)**2 + np.random.randn()/2) for i in range(self.size)])
            self.data = np.array(self.data)
        else:
            self.data = [(np.random.randn()/6, np.random.randn()/6) for i in range(150)]
            self.data.extend([(np.random.randn()/4 + 2.5, np.random.randn()/5) for i in range(150)])
            self.data.extend([(np.random.randn()/5 + 1, np.random.randn()/2 + 1) for i in range(150)])
            self.data.extend([(i/25 - 1, + np.random.randn()/20 - 1) for i in range(100)])
            self.data.extend([(i/25 - 2.5, 3 - (i/50 - 2)**2 + np.random.randn()/20) for i in range(150)])
            self.data = np.array(self.data)

    def print_data(self):
        print(self.data[:])

    def save_data_to_csv(self):
        import csv
        with open('data.csv', 'w+') as f:
            csv_writer = csv.writer(f, delimiter=',', quoting=csv.QUOTE_MINIMAL)
            for row in self.data:
                csv_writer.writerow(row)

    def plot_data(self):
        fig, ax = plt.subplots(figsize=(20, 10))
        ax.use_sticky_edges = False
        ax.margins(0