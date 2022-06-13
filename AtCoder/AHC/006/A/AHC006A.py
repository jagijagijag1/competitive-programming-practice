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
    寄り道をしたときに、
    最も距離の増加が少ないレストラン・配達先のペアを50個選ぶ
    """
    orders = []
    result_x = [CENTER, CENTER]
    result_y = [CENTER, CENTER]

    # 採用する注文の候補
    candidates = list(range(ORDER_COUNT))

    # オーダーの追加を50回繰り返す
    for _ in range(PICKUP_COUNT):
        min_dist = 1000000007
        # (選んだインデックス, レストラン挿入先, 配達先挿入先)
        min_indice = 0, 0, 0

        # l番目までを見たときの、最適なレストランの挿入位置を求めるための配列
        min_rest_dist = [1000000007] * len(result_x)
        min_rest_index = [998244353] * len(result_x)

        # 各オーダーについて、最適なレストラン・配達先の挿入先(j, k)を全探索
        for i, v in enumerate(candidates):
            x0 = rest_x[v]
            y0 = rest_y[v]
            x1 = dest_x[v]
            y1 = dest_y[v]

            min_rest_dist[0] = 1000000007

            # j番目までを見たときの、最適なレストランの挿入位置を求めておく
            for j in range(1, len(result_x)):
                prev_x = result_x[j - 1]
                prev_y = result_y[j - 1]
                next_x = result_x[j]
                next_y = result_y[j]

                old_dist_rest = dist(prev_x, prev_y, next_x, next_y)
                d00 = dist(prev_x, prev_y, x0, y0)
                d01 = dist(x0, y0, next_x, next_y)
                new_dist_rest = d00 + d01

                # j番目にレストランを挿入したときの距離の差分
                min_rest_dist[j] = new_dist_rest - old_dist_rest

                # 先頭から累積minを取っておく
                if min_rest_dist[j] < min_rest_dist[j - 1]:
                    min_rest_index[j] = j
                else:
                    min_rest_dist[j] = min_rest_dist[j - 1]
                    min_rest_index[j] = min_rest_index[j - 1]

            # 配達先の挿入先を探す
            for k in range(1, len(result_x)):
                prev_x = result_x[k - 1]
                prev_y = result_y[k - 1]
                next_x = result_x[k]
                next_y = result_y[k]

                # k - 1番目までにレストランを挿入したときの距離の差分
                dist_diff_rest = min_rest_dist[k - 1]

                old_dist_dest = dist(prev_x, prev_y, next_x, next_y)
                d10 = dist(prev_x, prev_y, x1, y1)
                d11 = dist(x1, y1, next_x, next_y)
                new_dist_dest = d10 + d11

                # k番目に配達先を挿入したときの距離の差分
                dist_diff_dest = new_dist_dest - old_dist_dest

                # (j, k)番目にそれぞれ挿入したときの距離の差分
                dist_diff = dist_diff_rest + dist_diff_dest

                # 最小値の更新
                if dist_diff < min_dist:
                    min_dist = dist_diff
                    # jを先に挿入するとkのインデックスが1つずれることに注意
                    min_indice = i, min_rest_index[k], k + 1

                # k番目にレストランと配達先を両方入れるケースもあるので、
                # この場合は場合分けが必要
                old_dist = dist(prev_x, prev_y, next_x, next_y)
                d10 = dist(prev_x, prev_y, x0, y0)
                d11 = dist(x0, y0, x1, y1)
                d12 = dist(x1, y1, next_x, next_y)
                new_dist = d10 + d11 + d12

                # (k, k)番目にそれぞれ挿入したときの距離の差分
                dist_diff = new_dist - old_dist

                # 最小値の更新
                if dist_diff < min_dist:
                    min_dist = dist_diff
                    # jを先に挿入するとkのインデックスが1つずれることに注意
                    min_indice = i, k, k + 1

        # オーダーの追加を行う
        i, j, k = min_indice
        v = candidates[i]
        orders.append(v)
        candidates.pop(i)

        # 寄り道する座標を挿入する。順番に注意
        result_x.insert(j, rest_x[v])
        result_y.insert(j, rest_y[v])
        result_x.insert(k, dest_x[v])
        result_y.insert(k, dest_y[v])

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
