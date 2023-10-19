const apiUrl :string = "https://dua-api.onrender.com/push-one"
import duaLipaImage from "@/types"
async function pushImagetoDB(duaLipaOBJ:duaLipaImage) {

    try {
        const res = await fetch(apiUrl,{
            method:"POST",
            body:JSON.stringify(
                {
                    "imgsrc":duaLipaOBJ.imgsrc,
                    "uploadedby":duaLipaOBJ.uploadedby,
                }
            )
        })
        const data = await res.json();
        console.log("POSTED ::",data);
    }catch(e:any) {
        console.log(e.message);
    }

}
export default pushImagetoDB