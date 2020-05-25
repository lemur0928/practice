import os
import requests
import json
from PIL import Image, ImageColor

url = "http://0.0.0.0:5000/model/predict"
image = '/Users/imlearning/Downloads/403466.jpg'#'/Users/imlearning/Downloads/S__41517085.jpg'
colors = {15: 'white', 9: 'red', 11: 'green'}
with open(image, 'rb') as im:
    fn = os.path.basename(image)
    files = {'image': (fn, im)}
    with requests.Session() as s:
        r = s.post(url, files=files)
        print(r)
        c = json.loads(r.content)

        im = Image.new('RGB', c['image_size']) 
        for (i,col) in enumerate(c['seg_map']):
            for (j,px) in enumerate(col):
                if px > 0:
                    print(px)
                    im.putpixel((j,i), ImageColor.getcolor(colors[px], 'RGB')) 
        im.save('segmentation.png')