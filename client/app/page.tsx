import Image from 'next/image'

export default function Home() {
  return (
    <center>
    <h1 style={{fontSize:20,marginTop:20}}>Worship The Goddess.</h1><br/>
    <input type="text" placeholder='Enter Dua Lipa Pic *Link*' accept='' style={{backgroundColor:"white",borderRadius:2}}/>&nbsp;&nbsp;&nbsp;
    <input type="text" placeholder='Your Name' style={{backgroundColor:"white",borderRadius:2}}/><br/><br/>
    <button style={{backgroundColor:"white",borderRadius:1.5}}>Upload!</button>
    <img src="https://hips.hearstapps.com/hmg-prod/images/gettyimages-1486911112-64503594c1449.jpg?crop=1xw:0.5xh;center,top&resize=1200:*" alt='Dua Pic' style={{maxWidth:"60%",margin:"20px"}} />
    <p>Uploaded By: Soubhik Gon</p>
    <img src="https://media.glamourmagazine.co.uk/photos/64b10f3eb673931dd09ca1b5/1:1/w_1280,h_1280,c_limit/DUA%20LIPA%20BARBIE%20140723%20default-GettyImages-1530547514.jpg" alt='Dua Pic' style={{maxWidth:"60%",margin:"20px"}} />
    <p>Uploaded By: Soubhik Gon</p>
    <img src="https://hips.hearstapps.com/hmg-prod/images/gettyimages-1486911112-64503594c1449.jpg?crop=1xw:0.5xh;center,top&resize=1200:*" alt='Dua Pic' style={{maxWidth:"60%",margin:"20px"}} />
    <p>Uploaded By: Soubhik Gon</p>
    </center>
    )
}