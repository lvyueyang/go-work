const { generateApi } = require('swagger-typescript-api');
const path = require('path');
const output = path.join(__dirname, './');
console.log('start generate api', output);
generateApi({
  name: 'index.ts',
  output,
  url: 'http://127.0.0.1:4002/swagger/doc.json',
  generateClient: false,
}).catch((e) => console.error(e));
