ORDER_COUNT = 1000  # 注文の数
PICKUP_COUNT = 50  # 選択する必要のある注文の数
CENTER = 400  # AtCoderオフィスのx,y座標


def read_input():
    """
    各点の座標を読み込む。
    """
    rest_x = []  # レストランのx座標
    rest_y = []  # レストランのy座標
    dest_x = []  # 配達先のx座標
    dest_y = []  # 配達先のy座標

    for _ in range(ORDER_COUNT):
        x0, y0, x1, y1 = map(int, input().split())
        rest_x.append(x0)
        rest_y.append(y0)
        dest_x.append(x1)
        dest_y.append(y1)

    return rest_x, rest_y, dest_x, dest_y


def solve(rest_x, rest_y, dest_x, dest_y):
    """
    今いる場所に最も近いレストランを50個貪欲に回り、
    その後最も近い配達先を貪欲に回る
    """
    orders = []
    result_x = [CENTER]
    result_y = [CENTER]

    # 採用する注文の候補（1～1000番の全てを候補とする）
    candidates = list(range(ORDER_COUNT))

    # 現在位置をオフィスにセット
    current_x = CENTER
    current_y = CENTER

    # 今いる場所に最も近いレストランを50個貪欲に回る
    for _ in range(PICKUP_COUNT):
        # 候補の中から、一番近いレストランを探す
        min_dist = 1000000007
        min_index = 998244353

        for i, v in enumerate(candidates):
            d = dist(current_x, current_y, rest_x[v], rest_y[v])
            if d < min_dist:
                min_dist = d
                min_index = i

        # レストランの注文番号
        v = candidates[min_index]

        # 解を更新
        orders.append(v)
        result_x.append(rest_x[v])
        result_y.append(rest_y[v])

        # 現在位置を更新
        current_x = rest_x[v]
        current_y = rest_y[v]

        # 既に到着したレストランは候補から削除
        candidates.pop(min_index)

    # 行かなきゃいけない配達先（＝これまでに行ったレストラン）のリスト
    candidates = orders.copy()

    # 今いる場所に最も近い配達先を貪欲に回る
    for _ in range(PICKUP_COUNT):
        # まだ行っていない配達先の中から、一番近い配達先を探す
        min_dist = 1000000007
        min_index = 998244353

        for i, v in enumerate(candidates):
            d = dist(current_x, current_y, dest_x[v], dest_y[v])
            if d < min_dist:
                min_dist = d
                min_index = i

        # 配達先の注文番号
        v = candidates[min_index]

        # 解を更新
        result_x.append(dest_x[v])
        result_y.append(dest_y[v])

        # 現在位置を更新
        current_x = dest_x[v]
        current_y = dest_y[v]

        # 既に到着した配達先は候補から削除
        candidates.pop(min_index)

    # 最後にオフィスに戻る
    result_x.append(CENTER)
    result_y.append(CENTER)

    return orders, result_x, result_y


def output(orders, x, y):
    """
    解を出力する。
    """
    orders = [str(i + 1) for i in orders]
    coordinates = [f"{x} {y}" for x, y in zip(x, y)]
    print(f"{len(orders)} {' '.join(orders)}")
    print(f"{len(coordinates)} {' '.join(coordinates)}")


def dist(x0, y0, x1, y1):
    """
    (x0, y0)と(x1, y1)のマンハッタン距離を求める
    """
    return abs(x0 - x1) + abs(y0 - y1)


rest_x, rest_y, dest_x, dest_y = read_input()
orders, x, y = solve(rest_x, rest_y, dest_x, dest_y)
output(orders, x, y)
