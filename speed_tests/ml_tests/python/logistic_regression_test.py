import numpy as np
from sklearn.linear_model import LogisticRegression
from common import *


class LogisticRegressionTester:

    X = None,
    y = None,
    X_test = None,
    y_test = None,
    logr = None

    def __init__(self):
        pass

    def set_data(self, path: str, path_test: str):
        train = np.genfromtxt(path, delimiter=',')[1:]
        train_test = np.genfromtxt(path_test, delimiter=',')[1:]
        self.X = train[:, 1:]
        self.y = train[:, 0]
        self.X_test = train_test[:, 1:]
        self.y_test = train_test[:, 0]
        print(self.X.shape)

    def create_linear_regression(self):
        self.logr = LogisticRegression(random_state=0, solver='lbfgs')
        self.do_logistic_