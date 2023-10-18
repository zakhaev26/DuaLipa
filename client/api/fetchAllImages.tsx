const apiUrl = "http://localhost:6969/all-image";

async function getImagesAPI() {
    try {
        const response = await fetch(apiUrl,{
            method:"GET"
            })
        const data = await response.json()
        return data;
    }catch(e:any) {
        console.log(e.message);
    }

}

export default getImagesAPI