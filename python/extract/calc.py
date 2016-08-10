#! /usr/bin/env python
# -*- coding: utf-8 -*-


def main():
    f = open("log.txt")
    read_mode = False
    # query_dict = {}
    while True:
        line = f.readline()
        if not line:
            break
        line = line.strip()
        # line = line.strip()
        if line == "{":
            stack = [True]
            json_str = line
            read_mode = True
            while read_mode == True:
                line = f.readline().strip()
                json_str += line
                if len(line) > 0 and line[len(line) - 1] == '}':
                    stack.pop()
                if len(line) > 1 and line[len(line) - 1] == ',':
                    if line[len(line) - 2] == '}':
                        stack.pop()
                if len(line) > 0 and line[len(line) - 1] == '{':
                    stack.append(True)
                if len(stack) == 0:
                    read_mode = False
            print json_str


if __name__ == '__main__':
    main()
