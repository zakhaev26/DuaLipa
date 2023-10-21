const apiUrl = "https://dua-api.onrender.com/push-one"
async function pushImagetoDB(duaLipaOBJ) {

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
    }catch(e) {
        console.log(e.message);
    }

}
export default pushImagetoDB