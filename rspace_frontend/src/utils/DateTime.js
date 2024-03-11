

// 将datetime类型的数据转为正常的时间显示：
// 例如：2024-03-08T15:10:56+08:00  ---->     
const FormatDateTime = (dateTimeString) =>{
    const dateTime = new Date(dateTimeString);
    
    const year = dateTime.getFullYear();
    const month = String(dateTime.getMonth() + 1).padStart(2, '0'); // 注意月份是从0开始的，所以要加1
    const day = String(dateTime.getDate()).padStart(2, '0');
    const hours = String(dateTime.getHours()).padStart(2, '0');
    const minutes = String(dateTime.getMinutes()).padStart(2, '0');
    const seconds = String(dateTime.getSeconds()).padStart(2, '0');

    const formattedDateTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    return formattedDateTime
}


export default FormatDateTime