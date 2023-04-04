const readline = require("readline");
const readObj = readline.createInterface({
  input: process.stdin
})

readObj.question('', (input) => {
//   inputArray = input.trim();
//   console.log(inputArray[0] + inputArray[1]);


  const input0 = input.trim();//inputArray[0];
  const input1 = input.trim();//inputArray[1];
  console.log(input0, input1);
  if(input0.length != input1.length) {
    console.log(0);
    readObj.close();
    return;
  } else {
      let obj = new Object();
      for (let index = 0; index < input0.length; index++) {
        const key = input0[index];
        const value = input1[index];
        if(obj[key] || obj[value]) {
            if(obj[key]!=value) {
                console.log(0);
                readObj.close();
                return;
            }
        } else {
            obj[key] = value;
            obj[value] = key;
        }
      }
  }
  console.log(1);


  readObj.close();
});
