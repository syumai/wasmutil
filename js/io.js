class ArrayReader {
  constructor(ary) {
    this.ary = ary;
    this.pos = 0;
  }

  get length() {
    return this.ary.length;
  }

  read(buf) {
    if (!buf) {
      return {
        eof: true,
        nread: 0,
      };
    }

    const restLen = this.length - this.pos;

    let nread, eof;
    if (buf.length > restLen) {
      nread = restLen;
      eof = true;
    } else {
      nread = buf.length;
      eof = false;
    }

    for (let i = 0; i < nread; i++) {
      buf[i] = this.ary[i + this.pos];
    }
    this.pos = this.pos + nread;

    return {
      nread,
      eof,
    };
  };
}
