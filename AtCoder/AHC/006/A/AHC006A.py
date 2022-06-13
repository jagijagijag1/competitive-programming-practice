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
    今いる場所に最も近いレストランまたは配達先を貪欲に回る
    ただし、候補はオフィスから距離400以下の場所のみとする
    """
    orders = []
    result_x = [CENTER]
    result_y = [CENTER]

    # 距離の閾値
    DEST_MAX = 400

    # 採用する注文の候補
    candidates_rest = []
    candidates_dest = []

    for i in range(ORDER_COUNT):
        dist_rest = dist(CENTER, CENTER, rest_x[i], rest_y[i])
        dist_dest = dist(CENTER, CENTER, dest_x[i], dest_y[i])
        if dist_rest <= DEST_MAX and dist_dest <= DEST_MAX:
            candidates_rest.append(i)

    current_x = CENTER
    current_y = CENTER

    # [NEW!] 今いる場所に最も近いレストランまたは配達先を貪欲に回る
    for _ in range(PICKUP_COUNT * 2):
        min_dist = 1000000007
        min_index = 998244353
        is_restaurant = True

        # レストランを探索
        for i, v in enumerate(candidates_rest):
            d = dist(current_x, current_y, rest_x[v], rest_y[v])
            if d < min_dist:
                min_dist = d
                min_index = i
                is_restaurant = True

        # 配達先を探索
        for i, v in enumerate(candidates_dest):
            d = dist(current_x, current_y, dest_x[v], dest_y[v])
            if d < min_dist:
                min_dist = d
                min_index = i
                is_restaurant = False

        if is_restaurant:
            # レストランに行く場合
            v = candidates_rest[min_index]
            x = rest_x[v]
            y = rest_y[v]
            orders.append(v)
            candidates_dest.append(v)
            if len(orders) >= PICKUP_COUNT:
                # 50個になったらそれ以上レストランを探さない
                candidates_rest.clear()
            else:
                candidates_rest.pop(min_index)
        else:
            # 配達先に行く場合
            v = candidates_dest[min_index]
            x = dest_x[v]
            y = dest_y[v]
            candidates_dest.pop(min_index)

        result_x.append(x)
        result_y.append(y)
        current_x = x
        current_y = y

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
