//replace with your own backend server ip
const ip = "http://172.24.190.128:9453"

export async function previewFile() {
    var file    = document.querySelector('input[type=file]').files[0];
    var reader  = new FileReader();

    var namespace = String(document.getElementsByClassName("input-namespace")[0].value)

    if (file && namespace!=="") {
      reader.readAsArrayBuffer(file);
      reader.onload = async function () {
        var data = reader.result;
        var buffer = data;
        var view = new Uint8Array(buffer);
        console.log(view)

        var res = await sendtoserver(view, namespace).then(res=>{
          return res.json()
        });
        console.log(res.blockheight)
        console.log(res.tx_hash)
        //download link create
        document.getElementsByClassName("submitresult")[0].innerText = "Blockheight: "+res.blockheight+"  Tx_hash: "+res.tx_hash
      };
    }
    
  }

export async function get_handler(){
  var namespace = String(document.getElementsByClassName("get-namespace")[0].value)
  var height = (document.getElementsByClassName("get-blockheight")[0].value)
  var res = await get_data(namespace, height).then(res=>{return res.json()})
  const blobData = atob(res.Blob);
  // Create a Blob object from the Base64-decoded data
  var reader  = new FileReader();
  const blob = new Blob([new Uint8Array([...blobData].map(char => char.charCodeAt(0)))], { type: 'application/octet-stream' });
  console.log(blob)
  reader.readAsArrayBuffer(blob);
    reader.onload = async function () {
      var data = reader.result;
      var buffer = data;
      var view = new Uint8Array(buffer);
      var preview = document.querySelector('img');
      console.log(view)
      preview.src = 'data:image/jpg;base64,' + tobase64encode(view);
      
    }
}

const sendtoserver = (bytearray, name) => {
  return new Promise((resolve, reject) => {
    const formData = new FormData();
    formData.append("arr", new Blob([bytearray], {type : ''}));
    formData.append("name", name);
    fetch(ip + "/submit", {
      method: 'POST',
      body: formData
  }).then(res=>{resolve(res)})
  });
}

const get_data = (name, height) => {
  return new Promise((resolve, reject) => {
    const formData = new FormData();
    formData.append("height", height);
    formData.append("name", name);
    fetch(ip + "/get", {
      method: 'POST',
      body: formData
  }).then(res=>{resolve(res)})
  });
}


function tobase64encode(input) {
  var keyStr =
          "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
  var output = "";
  var chr1, chr2, chr3, enc1, enc2, enc3, enc4;
  var i = 0;
 
  while (i < input.length) {
    chr1 = input[i++];
    chr2 = i < input.length ? input[i++] : Number.NaN;
    chr3 = i < input.length ? input[i++] : Number.NaN;
 
    enc1 = chr1 >> 2;
    enc2 = ((chr1 & 3) << 4) | (chr2 >> 4);
    enc3 = ((chr2 & 15) << 2) | (chr3 >> 6);
    enc4 = chr3 & 63;
 
    if (isNaN(chr2)) {
      enc3 = enc4 = 64;
    } else if (isNaN(chr3)) {
      enc4 = 64;
    }
    output += keyStr.charAt(enc1) + keyStr.charAt(enc2) + 
              keyStr.charAt(enc3) + keyStr.charAt(enc4);
  }
  return output;
}
