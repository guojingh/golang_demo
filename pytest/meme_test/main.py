import base64
import httpx


def main():
    with open("test.png", "rb") as f:
        base64_data = base64.b64encode(f.read()).decode()

    url = "http://127.0.0.1:2233/image/upload"
    payload = {"type": "data", "data": base64_data}
    resp = httpx.post(url, json=payload)
    image_id = resp.json()["image_id"]
    print(image_id)
    
    url = "http://127.0.0.1:2233/memes/petpet"
    payload = {
        "images": [{"name": "test", "id": image_id}],
        "texts": [],
        "options": {"circle": True},
    }
    resp = httpx.post(url, json=payload)
    image_id = resp.json()["image_id"]

    url = f"http://127.0.0.1:2233/image/{image_id}"
    resp = httpx.get(url)
    with open("result.gif", "wb") as f:
      f.write(resp.content)


if __name__ == "__main__":
    main()