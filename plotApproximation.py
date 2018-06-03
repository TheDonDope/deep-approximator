from mpl_toolkits.mplot3d import Axes3D
import matplotlib.pyplot as plt
import numpy as np
import math
import json
import sys
import os
import random

def draw3D():
    fig = plt.figure()
    coordinates = []
    with open("coordinates.json") as f:
        coordinates = json.load(f)


    ax = fig.add_subplot(111, projection='3d')

    for elem in coordinates:
        x = elem['x']
        y = elem['y']
        z1 = elem['z']
        z2 = math.sin(abs(math.sin(0.25*x*x + 0.5*y*y)))

        ax.scatter(x, y, z1, c='b', marker='x')
        ax.scatter(x, y, z2, c='g', marker='o')

        # error visualized via line, long=bad
        ax.plot([x, x], [y, y], zs=[z1, z2], c='r')

    ax.set_xlabel('X Label')
    ax.set_ylabel('Y Label')
    ax.set_zlabel('Z Label')

    plt.show()

def execCmd(cmd):
    print(cmd)
    os.system(cmd + " 2>/dev/null")

if __name__ == '__main__':

    rounds = 100000
    startSpeedReziprok = 2.
    endSpeedReziprok = 4.
    speedStep = 1.

    if sys.argv[1] == 'print':
        draw3D()
    elif sys.argv[1] == 'new':
        # delete old data
        cmd = "rm results.json"
        execCmd(cmd)
        # create network -> results.json
        cmd = "go run cmd/deep-approximator/main.go -l --hidden-layers 3 -n 100 -r 10 -s 0.1 -o results.json"
        execCmd(cmd)
        for x in np.arange(startSpeedReziprok, endSpeedReziprok, speedStep):
            # speed = 1./math.pow(2,x)
            speed = random.uniform(0.01, 0.33)
            # learn
            cmd = "go run cmd/deep-approximator/main.go -l -i results.json          -r " + \
                str(int(rounds)) + " -s " + str(speed) + " -o results.json"
            try:
                execCmd(cmd)
            except KeyboardInterrupt:
                print 'Interrupted'
                try:
                    sys.exit(0)
                except SystemExit:
                    os._exit(0)
            # calculate
            cmd = "go run cmd/deep-approximator/main.go -c -i results.json"
            execCmd(cmd)
        # draw 3d
        draw3D()
    elif sys.argv[1] == 'continue':
        for whatever in range(1, 100, 1):
            for x in np.arange(startSpeedReziprok, endSpeedReziprok, speedStep):
                speed = random.uniform(0.01, 0.33)
                # learn
                cmd = "go run cmd/deep-approximator/main.go -l -i results.json          -r " + \
                    str(int(rounds))+" -s " + str(speed) + " -o results.json"
                try:
                    execCmd(cmd)
                except KeyboardInterrupt:
                    print 'Interrupted'
                    try:
                        sys.exit(0)
                    except SystemExit:
                        os._exit(0)
                # calculate
                cmd = "go run cmd/deep-approximator/main.go -c -i results.json"
                execCmd(cmd)
        # draw 3d
        draw3D()
    else:
        print("Usage: python plotApproximation.py [full,print]")




