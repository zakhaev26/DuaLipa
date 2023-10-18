"use client"
import getImagesAPI from '@/api/fetchAllImages'
import {useEffect,useState} from "react";
import duaLipaImage from '@/types';
import pushImagetoDB from '@/api/pushImage';
export default function Home() {

  const [images,setImages] = useState([]);
  const [imgUrl,setImgUrl] = useState("");
  const [uploader,setUploader] = useState("");
  
  useEffect(()=>{
    async function callAPI() {
      try{
        const res = await getImagesAPI();
        console.log("Done",res);
        setImages(res);
      }catch(e:any) {
        console.log(e.message);
      }
    }
    callAPI();
  },[]);

  function handleChange(event :any,setter:any) {
      setter(event.target.value);
  }

  async function handleClick() {
    if (!imgUrl || !uploader) {
      alert("Please Fill up the fields Correctly and then Submit.")
      return;
    }
    let duaLipaOBJ : duaLipaImage ={
      imgsrc : imgUrl,
      uploadedby:uploader,
    }
    try{
      const res = pushImagetoDB(duaLipaOBJ);
      console.log("Posted::",res)
      alert("Posted!");
      setImgUrl("");
      setUploader("");
    }catch(e:any) {
      console.log(e.message);
    }
  }
  return (
    <center>
    <h1 style={{fontSize:20,marginTop:20}}>Worship The Goddess.</h1><br/>
    <input onChange={(e)=>handleChange(e,setImgUrl)} value={imgUrl} type="text" placeholder='Enter Dua Lipa Pic *Link*' accept='' style={{backgroundColor:"white",borderRadius:2}}/>&nbsp;&nbsp;&nbsp;
    <input onChange={(e)=>handleChange(e,setUploader)} type="text" placeholder='Your Name' style={{backgroundColor:"white",borderRadius:2}}/><br/><br/>
    <button onClick={handleClick} style={{backgroundColor:"white",borderRadius:1.5}}>Upload!</button>
    {
      images.map((image:duaLipaImage,index:number)=>(
        <div>
        <img key={index} src={image.imgsrc} alt='Dua Pic' style={{maxWidth:"60%",margin:"20px"}} />
        <p>Uploaded By: {image.uploadedby}</p>
        </div>
      ))
    }
    </center>
    )
}