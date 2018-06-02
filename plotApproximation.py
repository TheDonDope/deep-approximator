from mpl_toolkits.mplot3d import Axes3D
import matplotlib.pyplot as plt
import numpy as np
import math
import json
import sys

coordinates = []
with open("coordinates.json") as f:
    coordinates = json.load(f)

# for x in np.arange(-1.0, 1.0, 0.2):
#     for y in np.arange(-1.0, 1.0, 0.2):
#         coordinates.append([x, y, math.sin(x + y)])

fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')

for elem in coordinates:
    x = elem['x']
    y = elem['y']
    z1 = elem['z']
    z2 = math.sin(x*x + y)
    # z2 = math.sin(x-y)
    ax.scatter(x, y, z1, c='b', marker='.')
    ax.scatter(x, y, z2, c='r', marker='*')
    ax.plot([x, x], [y, y], zs=[z1, z2], c='b')

ax.set_xlabel('X Label')
ax.set_ylabel('Y Label')
ax.set_zlabel('Z Label')

plt.show()
