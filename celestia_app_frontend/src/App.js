import './App.css';
import { previewFile, get_handler } from './loadfile';

function App() {
  return (
    <div className="App">
        <div className="container">
        <div className="content">
        <form name='submitform' action='#'>
          <div className="topic">Submit Photo</div>
          <input type="file" accept=".jpg" /><br />
          <div className="input-box">
            <input className="input-namespace" type="text" required />
            <label>namespace</label>
          </div>
          <div className="input-box">
            <div className='submitresult'></div>
            <div className='submitresult_txhash'></div>
          </div>
          <div className="input-box">
            <input type="submit" value="Submit" onClick={previewFile}/>
          </div>
        </form>
        <form action="#">
          <div className="topic">Get Photo</div>
          <div className="input-box">
            <input className="get-namespace" type="text" required />
            <label>namespace</label>
          </div>
          <div className="input-box">
            <input className="get-blockheight" type="text" required />
            <label>blockheight</label>
          </div>
          <div className="input-box">
            <input type="submit" value="Get" onClick={get_handler}/>
          </div>
          <div className="image-box">
            <img/>
          </div>
        </form>
      </div>
      </div>
    </div>
  );
}

export default App;
