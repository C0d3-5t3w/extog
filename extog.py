import logging
import pwnagotchi.plugins as plugins
import subprocess
from flask import render_template_string

class ExternalWiFiToggle(plugins.Plugin):
    __author__ = '5T3W'
    __version__ = '1.0.0'
    __license__ = 'MIT'
    __description__ = 'A plugin to toggle between internal and external WiFi adapter, must be used with extog application.'

    def __init__(self):
        self.ready = False

    def on_loaded(self):
        logging.info("[ExternalWiFiToggle] plugin loaded.")
        self.ready = True

    def on_webhook(self, path, request):
        if not self.ready:
            return "Plugin not ready"

        logging.info(f"[ExternalWiFiToggle] DEBUG - Full request: {request}")
        logging.info(f"[ExternalWiFiToggle] DEBUG - Path: '{path}'")
        logging.info(f"[ExternalWiFiToggle] DEBUG - Method: {request.method}")
        logging.info(f"[ExternalWiFiToggle] DEBUG - Form data: {request.form}")
        logging.info(f"[ExternalWiFiToggle] DEBUG - Args: {request.args}")

        try:
            if request.method == "GET":
                ret = '''
                <html>
                    <head>
                        <title>WiFi Adapter Toggle</title>
                        <meta name="csrf_token" content="{{ csrf_token() }}">
                    </head>
                    <body>
                        <h1>WiFi Adapter Toggle</h1>
                        <form method="POST" action="">
                            <input type="hidden" name="csrf_token" value="{{ csrf_token() }}">
                            <input type="submit" name="state" value="ON" style="font-size: 20px; padding: 10px 20px; margin: 5px;">
                            <input type="submit" name="state" value="OFF" style="font-size: 20px; padding: 10px 20px; margin: 5px;">
                        </form>
                        <p>ON = Use External WiFi Adapter<br>OFF = Use Internal WiFi Adapter</p>
                    </body>
                </html>
                '''
                return render_template_string(ret)

            elif request.method == "POST":
                logging.info("[ExternalWiFiToggle] Processing POST request")
                state = request.form.get('state', '').lower()
                logging.info(f"[ExternalWiFiToggle] State value: {state}")
                
                if state in ['on', 'off']:
                    result = self._toggle_adapter(state)
                    ret = f'''
                    <html>
                        <head>
                            <title>WiFi Adapter Toggle Result</title>
                            <meta http-equiv="refresh" content="3;url=/plugins/ExternalWiFiToggle/">
                        </head>
                        <body>
                            <h1>{result}</h1>
                            <p>Redirecting back in 3 seconds...</p>
                        </body>
                    </html>
                    '''
                    return render_template_string(ret)
                else:
                    logging.error(f"[ExternalWiFiToggle] Invalid state value: {state}")
                    return "Invalid state value"

        except Exception as e:
            logging.error(f"[ExternalWiFiToggle] Web Error: {str(e)}")
            return f"Error: {str(e)}"

        return "Invalid request"

    def _toggle_adapter(self, state):
        try:
            logging.info(f"[ExternalWiFiToggle] Attempting to run extog with state: {state}")
            
            process = subprocess.Popen(['extog'], 
                                    stdin=subprocess.PIPE, 
                                    stdout=subprocess.PIPE,
                                    stderr=subprocess.PIPE,
                                    universal_newlines=True)
            
            process.stdin.write("yes\n")
            process.stdin.flush()
            
            process.stdin.write(f"{state}\n")
            process.stdin.flush()
            
            stdout, stderr = process.communicate()
            logging.info(f"[ExternalWiFiToggle] Output: {stdout}")
            if stderr:
                logging.error(f"[ExternalWiFiToggle] Errors: {stderr}")
            
            if process.returncode != 0:
                raise Exception(f"Process failed with return code {process.returncode}")
            
            return f"WiFi adapter switched to {'external' if state == 'on' else state == 'off'} successfully"
        except Exception as e:
            logging.error(f"[ExternalWiFiToggle] Error: {str(e)}")
            return f"Error toggling WiFi adapter: {str(e)}"
