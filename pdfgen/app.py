import json
import random
import time

import requests
from flask import Flask, send_file
from flask import request
app = Flask(__name__)

from PyPDF2 import  PdfWriter, PdfReader
from reportlab.pdfgen import canvas
from reportlab.lib.pagesizes import letter
from io import BytesIO


from reportlab.pdfbase import pdfmetrics
from reportlab.pdfbase.ttfonts import TTFont
pdfmetrics.registerFont(TTFont('Song', 'simsun.ttf'))


@app.route('/gen/', methods=['POST'])
def gen_pdf():
    id = request.json["id"]
    input_pdf = PdfReader(open('template.pdf', 'rb'))
    pdf_writer = PdfWriter()
    packet = BytesIO()
    can = canvas.Canvas(packet, pagesize=letter)
    can.setFont('Song', 12)
    can.drawString(90, 704, request.json["name"])
    can.drawString(200, 704, request.json["school"])
    can.drawString(310, 704, request.json["sid"])
    can.drawString(435, 704, request.json["contact"])
    can.drawString(150, 685, request.json["from"])
    can.drawString(435, 685, request.json["to"])
    can.setFont('Song', 12)
    can.drawString(50, 630, request.json["text0"])
    can.drawString(50, 535, request.json["text1"])
    can.drawString(50, 440, request.json["text2"])
    can.drawString(50, 350, request.json["text3"])
    can.drawString(50, 255, request.json["text4"])
    can.save()

    packet.seek(0)
    new_pdf = PdfReader(packet)
    page = input_pdf.pages[0]
    page.merge_page(new_pdf.pages[0])
    pdf_writer.add_page(page)

    with open(id+".pdf", 'wb') as out:
        pdf_writer.write(out)

    res = requests.get("http://localhost:8080/api/v1/internal/pdf?id="+id)
    print(res.content)

    return str(id)


if __name__ == '__main__':
    app.run()