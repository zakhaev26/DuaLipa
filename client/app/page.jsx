"use client";
import getImagesAPI from "@/api/fetchAllImages";
import { useEffect, useState } from "react";
import pushImagetoDB from "@/api/pushImage";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Footer from "@/components/Footer";
export default function Home() {
  const notify = (message) => toast(message);
  const [images, setImages] = useState([]);
  const [imgUrl, setImgUrl] = useState("");
  const [uploader, setUploader] = useState("");

  useEffect(() => {
    async function callAPI() {
      try {
        const res = await getImagesAPI();
        console.log("Done", res);
        setImages(res);
      } catch (e) {
        console.log(e.message);
      }
    }
    callAPI();
  }, []);

  function handleChange(event, setter) {
    setter(event.target.value);
  }

  async function handleClick() {
    if (!imgUrl || !uploader) {
      // alert("Please Fill up the fields Correctly and then Submit.")
      notify("Please Fill up the fields Correctly and then Submit.");
      return;
    }
    let duaLipaOBJ = {
      imgsrc: imgUrl,
      uploadedby: uploader,
    };
    try {
      const res = pushImagetoDB(duaLipaOBJ);
      console.log("Posted::", res);
      notify("Posted.Refresh!");
      setImgUrl("");
      setUploader("");
    } catch (e) {
      console.log(e.message);
    }
  }
  return (
    <center>
      <img
        src="https://i.pinimg.com/originals/99/eb/d0/99ebd098906e1132efc9584852cee95c.gif"
        alt=""
      />
      <h1 style={{ fontSize: 20, marginTop: 20 }}>Worship The Goddess.</h1>
      <br />
      <input
        onChange={(e) => handleChange(e, setImgUrl)}
        value={imgUrl}
        type="text"
        placeholder="Enter Dua Lipa Pic *Link*"
        accept=""
        style={{ backgroundColor: "white", borderRadius: 2 }}
      />
<br />
<br />
      <input
        onChange={(e) => handleChange(e, setUploader)}
        type="text"
        placeholder="Your Name"
        style={{ backgroundColor: "white", borderRadius: 2 }}
      />
      <br />
      <br />
      <button
        onClick={handleClick}
        style={{ backgroundColor: "grey", borderRadius: 1.5 }}
      >
        Upload!
      </button>
      <ToastContainer />
      {images.map((image, index) => (
        <div key={index}>
          <img
            src={image.imgsrc}
            alt="Dua Pic"
            style={{ maxWidth: "50%", maxHeight: "300px", margin: "20px" }}
          />
          <p>Uploaded By: {image.uploadedby}</p>
        </div>
      ))}
      <br />
      <Footer />
    </center>
  );
}
