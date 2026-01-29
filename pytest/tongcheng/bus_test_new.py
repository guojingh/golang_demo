import time
import hashlib
import requests
import json
import sys  # 导入sys模块处理命令行参数
 
 
def generate_sign(params, token):
    """
    生成请求签名
    :param params: 请求参数字典
    :param token: 签名令牌
    :return: 签名值
    """
    # 移除sign参数（如果存在）
    params_to_sign = {k: v for k, v in params.items() if k != "sign"}
 
    # 按键名进行字典排序
    sorted_keys = sorted(params_to_sign.keys())
 
    # 构建待签名字符串
    sign_str_parts = []
    for key in sorted_keys:
        value = str(params_to_sign[key])
        sign_str_parts.append(f"{key}={value}")
    sign_param_str = "&".join(sign_str_parts)
 
    # 构建完整签名字符串
    timestamp = str(params["timeStamp"])
    app_id = str(params["appId"])
    full_sign_str = timestamp + app_id + token + sign_param_str
 
    # 调试信息
    print(f"待签名字符串: {full_sign_str}")
 
    # 计算MD5哈希
    return hashlib.md5(full_sign_str.encode('utf-8')).hexdigest()
 
 
def send_request(arrival, departure):
    # 基础参数
    app_id = "alamcp"
    token = "6ab065667d0564210a16e7f7e598ec4a"
    url = "http://10.41.108.145:8104/ticket/tapi/v1/bus_p2p_data"
 
    # 构造请求参数（从命令行获取arrival和departure）
    params = {
        "appId": app_id,
        "timeStamp": str(int(time.time())),  # 当前时间戳
        "dataType": "bus",
        "arrStation": arrival,  # 从命令行传入的目的地
        "depStation": departure,  # 从命令行传入的出发地
        "channel": "baikan"  # 渠道
    }
 
    # 生成签名并添加到参数
    sign = generate_sign(params, token)
    params["sign"] = sign
 
    # 打印完整请求参数
    print("完整请求参数:")
    print(json.dumps(params, indent=2, ensure_ascii=False))
 
    # 设置请求头（匹配您的curl）
    headers = {
        "User-Agent": "Apifox/1.0.0 (https://apifox.com)",
        "Content-Type": "application/json",
        "Accept": "*/*",
        "Connection": "keep-alive"
    }
 
    try:
        # 发送JSON格式请求
        response = requests.post(
            url,
            json=params,  # 使用json参数自动序列化
            headers=headers,
            timeout=10
        )
 
        print(f"\n响应状态码: {response.status_code}")
        print(f"响应内容: {response.text}")
 
        # 尝试解析JSON响应
        try:
            return response.json()
        except json.JSONDecodeError:
            return response.text
    except requests.exceptions.RequestException as e:
        print(f"请求失败: {str(e)}")
        return None
 
 
if __name__ == "__main__":
    # 检查命令行参数数量
    if len(sys.argv) < 3:
        print("用法: python3 cartest.py <目的地> <出发地>")
        print("示例: python3 cartest.py 天津 北京")
        sys.exit(1)  # 退出程序，返回错误码
 
    # 获取命令行参数
    arrival = sys.argv[1]
    departure = sys.argv[2]
 
    print("开始发送请求...")
    print(f"出发地: {departure}, 目的地: {arrival}")
    result = send_request(arrival, departure)
 
    if result:
        print("\n请求结果:")
        if isinstance(result, dict):
            print(json.dumps(result, indent=2, ensure_ascii=False))
        else:
            print(result)