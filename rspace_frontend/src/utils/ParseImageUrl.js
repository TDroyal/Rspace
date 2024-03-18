
import BackendRootURL from '../common_resources/resource'

const prefix = BackendRootURL + "/static/posts/"


const ParseImageUrl = (imgstr)=>{
    let imageurl = []

    //对imgstr进行分割  "@#@"
    let len = imgstr.length

    for(let i = 0, start_idx = 0, end_idx = 0, id = 1; i < len; i ++ ) {
        if(imgstr[i] === '@' && imgstr[i + 1] === '#' && imgstr[i + 2] === '@') {
            end_idx = i
            imageurl.push({
                "id": id,
                "image":prefix + imgstr.substring(start_idx, end_idx),
            })
            start_idx = i + 3
            id ++ 
        }
        
    }
    return imageurl
}

export default ParseImageUrl