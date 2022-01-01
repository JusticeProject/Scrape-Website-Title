from flask import Flask, request, make_response

from gzip import GzipFile
import zlib
from io import BytesIO
import brotli

app = Flask(__name__)

###############################################################################

@app.route("/")
def index():
    html = '<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><title>Your Headers</title></head><body>'

    for key, value in request.headers:
        html += "{} = {}<br>".format(key, value)

    html += '<br><a href="/gzip">Click for gzip compressed page</a><br>'
    html += '<br><a href="/deflate">Click for deflate compressed page</a><br>'
    html += '<br><a href="/br">Click for br compressed page</a><br>'
    html += "</body></html>"
    return html

###############################################################################

@app.route("/gzip")
def gzip():
    html = '<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><title>Your Gzip Data</title></head><body>'
    html += "This page was compressed with gzip"
    html += "</body></html>"

    data = bytes(html, encoding="utf-8")
    gzip_buffer = BytesIO()
    with GzipFile(mode='wb', compresslevel=6, fileobj=gzip_buffer) as gzip_file:
        gzip_file.write(data)
    
    resp = make_response()
    resp.set_data(gzip_buffer.getvalue())
    resp.content_encoding = "gzip"

    return resp

###############################################################################

@app.route("/deflate")
def deflate():
    html = '<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><title>Your Deflate Data</title></head><body>'
    html += "This page was compressed with deflate"
    html += "</body></html>"

    data = bytes(html, encoding="utf-8")
    compressed_data = zlib.compress(data, -1)
    
    resp = make_response()
    resp.set_data(compressed_data)
    resp.content_encoding = "deflate"

    return resp

###############################################################################

@app.route("/br")
def br():
    html = '<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><title>Your Brotli Data</title></head><body>'
    html += "This page was compressed with brotli"
    html += "</body></html>"

    data = bytes(html, encoding="utf-8")
    compressed_data = brotli.compress(data, mode=0, quality=4, lgwin=22, lgblock=0)

    resp = make_response()
    resp.set_data(compressed_data)
    resp.content_encoding = "br"

    return resp

###############################################################################

if __name__ == "__main__":
    app.run("0.0.0.0", 6512)
