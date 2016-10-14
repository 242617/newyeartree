#include <Adafruit_NeoPixel.h>
#include <stdarg.h>

#define NUMPIXELS 4

Adafruit_NeoPixel pixels = Adafruit_NeoPixel(NUMPIXELS, 3, NEO_GRB + NEO_KHZ800);

void setup() {
  Serial.begin(115200);
  
  pixels.begin();
}

void loop() {
  if (Serial.available() >= 4) {
    int n = Serial.read();
    int r = Serial.read();
    int g = Serial.read();
    int b = Serial.read();

    pixels.setPixelColor(n, pixels.Color(r, g, b));
    pixels.show();

    p("number: %d, red: %d, green: %d, blue: %d\n", n, r, g, b);
  }
}

void p(char *fmt, ... ){
        char buf[128];
        va_list args;
        va_start (args, fmt );
        vsnprintf(buf, 128, fmt, args);
        va_end (args);
        Serial.print(buf);
}

