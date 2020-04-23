#version 400

uniform vec3 color;
uniform float scale;
uniform float height;
uniform float width;

out vec4 frag;

void main() {
  vec2 pos = gl_FragCoord.xy / scale;

  vec3 white = vec3(1., 1., 1.);

  frag = vec4(mix(white, color, pos.x / width), 1.);
}