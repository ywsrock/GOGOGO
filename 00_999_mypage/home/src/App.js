import './App.css';
import Nav from './component/Nav';
import { useEffect } from 'react';
import { Outlet, useNavigation } from 'react-router-dom';

// const bgImage = require("./assets/bg.png")
const bgImage = require("./assets/gopher.png")
const style = {
  backgroundImage: `url(${bgImage})`,
  backgroundRepeat: "no-repeat",
  backgroundPosition: "center",
  backgroundSize: "cover",
}

function App() {
  // eslint-disable-next-line no-undef
  useEffect(() => {
    document.title = 'yws-page'
  }, [])

  // Golbal Pending Ui
  const navigation = useNavigation();

  return (
    <div div style={style} >
      <Nav></Nav>
      <div id='main-page' className={navigation.state === "loading" ? "loading" : ""
      }>
        <Outlet />
      </div>
    </div >
  );
}
export default App;
