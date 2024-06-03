import {useEffect,useState} from "react"
import './App.css';
import { connect,senMsg } from './api';

function App() {
  const [inputValue, setInputValue] = useState('');
  useEffect(() => {
    connect();
  }, [])

  function send(){
    connect();
    console.log(inputValue);
    senMsg(inputValue)
    setInputValue('')
  }

  return (
    <>
    <div className="App">
      <input type="text" value={inputValue} onChange={(event) => setInputValue(event.target.value)} />
      <button onClick={send}>hit</button>
    </div>
    <p id="output"></p>
    </>
  );
}

export default App;
