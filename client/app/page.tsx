"use client"
import getImagesAPI from '@/api/fetchAllImages'
import {useEffect,useState} from "react";
import duaLipaImage from '@/types';
export default function Home() {

  const [images,setImages] = useState([]);

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

  return (
    <center>
    <h1 style={{fontSize:20,marginTop:20}}>Worship The Goddess.</h1><br/>
    <input type="text" placeholder='Enter Dua Lipa Pic *Link*' accept='' style={{backgroundColor:"white",borderRadius:2}}/>&nbsp;&nbsp;&nbsp;
    <input type="text" placeholder='Your Name' style={{backgroundColor:"white",borderRadius:2}}/><br/><br/>
    <button style={{backgroundColor:"white",borderRadius:1.5}}>Upload!</button>
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