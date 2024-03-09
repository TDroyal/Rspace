

// 将datetime类型的数据转为正常的时间显示：
// 例如：2024-03-08T15:10:56+08:00  ---->     
const FormatDateTime = (dateTimeString) =>{
    const dateTime = new Date(dateTimeString);

    const year = dateTime.getFullYear();
    const month = dateTime.getMonth() + 1; // 注意月份是从0开始的，所以要加1
    const day = dateTime.getDate();
    const hours = dateTime.getHours();
    const minutes = dateTime.getMinutes();
    const seconds = dateTime.getSeconds();

    const formattedDateTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    return formattedDateTime
}


export default FormatDateTime