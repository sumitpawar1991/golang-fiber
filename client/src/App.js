import {Routes,Route,BrowserRouter} from "react-router-dom";

import "./App.css"
import Home from "./page/Home";
import Blog from "./page/Blog";
import Header from "./components/layout/Header";
import Footer from "./components/layout/Footer";
import About from "./page/About";
import Contact from "./page/Contact";
import BlogDetail from "./page/BlogDetail";
import Signin from "./page/Signin";
import Singup from "./page/Singup";



function App() {
return (
  <BrowserRouter>
      <Header/>

      <main style={{ padding:"20px"}}>
        <Routes>

          <Route path="/" element={<Home />}/>
          <Route path="/blog" element={<Blog />}/>
          <Route path="/blog/:id" element={<BlogDetail />}/>
          <Route path="/about" element={<About />}/>
          <Route path="/contact" element={<Contact />}/>
          <Route path="/signin" element={<Signin />}/>
          <Route path="/signup" element={<Singup />}/>
        </Routes>
      </main>

      <Footer/>
  </BrowserRouter>
  );
}

export default App;
