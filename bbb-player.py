import argparse
from urllib.parse import urlparse
import os
import urllib.request
import json
from distutils.dir_util import copy_tree
import traceback
import re
from datetime import timedelta
import logging
import webbrowser
from sys import exit



SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))
logging.basicConfig(format="[%(asctime)s -%(levelname)8s]: %(message)s",
                    datefmt="%H:%M:%S",
                    level=logging.INFO)
logger = logging.getLogger('bbb-player')

logger.setLevel(logging.WARNING)

try:
    from flask import Flask, render_template, request, redirect, url_for
except:
    logger.error("Flask not imported. Try running:\npip3 install Flask")
    exit(1)


logger.warning('---------------------------')
logger.warning('Dokuz Eylul University    ')
logger.warning('BigBlueButton Recording Viewer')
logger.warning('Author: github.com/maaami98')

logger.warning('---------------------------')
logger.warning('In your modern web browser open:')
logger.warning('http://localhost:5000')
logger.warning('Press CTRL+C when done.')
logger.warning('----------------------')


# Based on https://stackoverflow.com/a/42791810
# Flask is needed for HTTP 206 Partial Content support.
app = Flask(__name__)
app.config["DEBUG"] = False

@app.route("/", methods=["GET"])
def hello():
    
    return redirect('/public/index.html',code=302)
# Based on https://stackoverflow.com/a/37331139
# This is needed for playback of multiple meetings in short succession.
app.config['SEND_FILE_MAX_AGE_DEFAULT'] = 0
app.config['TESTING'] = False
app.config["DEBUG"] = False
import webbrowser
if __name__ == "__main__":
    webbrowser.open("http://127.0.0.1:5000")
    app.run(host='127.0.0.1', port=5000) # Securty 127.0.0.1
    
