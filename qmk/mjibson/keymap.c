// this is the style you want to emulate.
// This is the canonical layout file for the Quantum project. If you want to add another keyboard,
#include QMK_KEYBOARD_H
// Each layer gets a name for readability, which is then used in the keymap matrix below.
// The underscores don't mean anything - you can have a layer called STUFF or any other name.
// Layer names don't all need to be of the same length, obviously, and you can also skip them
// entirely and just use numbers.
enum ferris_layers {
  _DVORAK,
  _LOWER,
  _PLOVER,
};
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
  [_DVORAK] = LAYOUT( /* DVORAK */
    KC_QUOT,        KC_COMM,        KC_DOT,          KC_P,           KC_Y,           KC_F,            KC_G,            KC_C,            KC_R,           KC_L,
    KC_A,           KC_O,           KC_E,            KC_U,           KC_I,           KC_D,            KC_H,            KC_T,            KC_N,           KC_S,
    LT(1, KC_SCLN), LCTL_T(KC_Q),   LALT_T(KC_J),    LCMD_T(KC_K),   KC_X,           KC_B,            LCMD_T(KC_M),    LALT_T(KC_W),    LCTL_T(KC_V),   LT(1, KC_Z),
                                                     LT(1, KC_TAB),  LSFT_T(KC_SPC), LSFT_T(KC_ENT),  LT(1, KC_BSPC)
  ),
  [_LOWER] = LAYOUT( /* [> LOWER <] */
    KC_1,           KC_2,           KC_3,            KC_4,           KC_5,           KC_6,            KC_7,            KC_8,            KC_9,            KC_0,
    KC_GRAVE,       KC_SLSH,        KC_LBRC,         KC_RBRC,        KC_INS,         KC_VOLU,         KC_LEFT,         KC_DOWN,         KC_UP,           KC_RGHT,
    KC_BSLS,        LCTL_T(KC_EQL), LALT_T(KC_MINS), LCMD_T(KC_ESC), KC_DEL,         KC_VOLD,         LCMD_T(KC_HOME), LALT_T(KC_PGDN), LCTL_T(KC_PGUP), KC_END,
                                                     LT(1, KC_TAB),  LSFT_T(KC_SPC), LSFT_T(KC_MUTE), LT(1, KC_MPLY)
  ),
  [_PLOVER] = LAYOUT( /* [> PLOVER <] */
    KC_1,    KC_1,    KC_1,    KC_1,    KC_1,    KC_1,    KC_1,    KC_1,    KC_1,    TO(0),
    KC_Q,    KC_W,    KC_E,    KC_R,    KC_T,    KC_U,    KC_I,    KC_O,    KC_P,    KC_LBRC,
    KC_A,    KC_S,    KC_D,    KC_F,    KC_G,    KC_J,    KC_K,    KC_L,    KC_SCLN, KC_QUOT,
                               KC_C,    KC_V,    KC_N,    KC_M
  ),
};
