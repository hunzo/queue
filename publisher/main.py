from loads.services import tasks, tasks_queue


def main():
    print("start")
    test_list = [
        # 1000,
        # 1000,
        # 1000,
        # 1000,
        # 1000,
        # 1000,
        # 1000,
        # 1000,
        # 1000,
        # 1000,
        # 1000,
        # 100000,
        10000,
    ]


    for i in test_list:
        tasks_queue(i)
        # tasks(i)
    
    pass

if __name__ == "__main__":
    main()
