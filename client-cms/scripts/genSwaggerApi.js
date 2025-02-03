const { generateApi } = require('swagger-typescript-api');
const path = require('path');

generateApi({
  name: 'serverApi.ts',
  output: path.join(__dirname, '../src/interface'),
  url: 'http://127.0.0.1:4002/swagger/doc.json',
  generateClient: false,
}).catch((e) => console.error(e));
