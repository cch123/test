import datetime
import sys

def main():
    if len(sys.argv[1:]) == 0:
        print "Param number error\n Pleaserun script likt python xxx.py 20150432"
        sys.exit()

    dt = datetime.datetime.strptime(sys.argv[1], "%Y%m%d")
    print dt

if __name__ == "__main__":
    main()
