meta:
  id: dc6
  title: Diablo CEL 6
  application: Diablo II
  file-extension: dc6
  license: MIT
  ks-version: 0.7
  encoding: ASCII
  endian: le
seq:
  - id: dc6
    type: file
types:
  file:
    seq:
      - id: header
        type: file_header
      - id: frame_pointers
        type: u4
        repeat: expr
        repeat-expr: header.directions * header.frames_per_dir
      - id: frames
        type: frame
        repeat: expr
        repeat-expr: header.directions * header.frames_per_dir
  file_header:
    seq:
      - id: version
        type: s4
      - id: flags
        type: u4
        enum: flags
      - id: encoding
        type: u4
      - id: termination
        size: 4
      - id: directions
        type: s4
      - id: frames_per_dir
        type: s4
    enums:
      flags:
        1: celfile_serialised
        4: celfile_24bit
  frame:
    seq:
      - id: header
        type: frame_header
      - id: block
        type: u1
        repeat: expr
        repeat-expr: header.length
      - id: terminator
        size: 3
    types:
      frame_header:
        seq:
          - id: flipped
            type: s4
          - id: width
            type: s4
          - id: height
            type: s4
          - id: offset_x
            type: s4
          - id: offset_y
            type: s4
          - id: unknown
            type: u4
          - id: next_block
            type: s4
          - id: length
            type: s4
