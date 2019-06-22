Blob.prototype.read = async function(buf) {
  if (!buf) {
    return {
      eof: true,
      nread: 0,
    };
  }
  const pos = this.pos || 0;
  const { size: thisLen } = this;
  const restLen = thisLen - pos;

  const { length: bufLen } = buf;
  let b, nread, eof;
  if (bufLen > restLen) {
    b = this.slice(pos, thisLen);
    nread = restLen;
    eof = true;
  } else {
    b = this.slice(pos, pos + bufLen);
    nread = bufLen;
    eof = false;
  }

  const fileReader = new FileReader();
  const p = new Promise(resolve => {
    fileReader.addEventListener('load', () => {
      resolve();
    });
  });
  fileReader.readAsArrayBuffer(b);
  await p;

  const data = new Uint8Array(fileReader.result);
  for (let i = 0; i < nread; i++) {
    buf[i] = data[i];
  }
  this.pos = pos + nread;

  return {
    nread,
    eof,
  };
};
