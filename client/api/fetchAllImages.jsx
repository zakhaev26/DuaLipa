const apiUrl = "https://dua-api.onrender.com/all-image";

async function getImagesAPI() {
    try {
        const response = await fetch(apiUrl,{
            method:"GET"
            })
        const data = await response.json()
        return data;
    }catch(e) {
        console.log(e.message);
    }

}

export default getImagesAPI