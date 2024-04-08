const max = (a, b)=>{
    if(a > b) return a;
    return b;
}


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

//获取时间所在的期间
const GetTimePeriod = (timeString) =>{
    // 将时间字符串转换为 Date 对象
    const time = new Date(timeString);
  
    // 获取当前时间
    const now = new Date();
  
    // 计算时间差（单位：毫秒）
    const diff = now - time;
  
    // 计算时间差的分钟数、小时数和天数
    const seconds = Math.floor(diff / 1000)
    const minutes = Math.floor(diff / (1000 * 60));
    const hours = Math.floor(diff / (1000 * 60 * 60));
    const days = Math.floor(diff / (1000 * 60 * 60 * 24));
    const months = Math.floor(diff / (1000 * 60 * 60 * 24 * 30))
  
    // 根据时间差返回相应的时间段字符串
    if (seconds <= 59) {
        return  max(seconds, 1) + "秒前"
    } else if (minutes <= 59) {
        return minutes + "分钟前";
    } else if (hours <= 23) {
        return hours + "小时前";
    } else if (days <= 30) {
        return days + "天前";
    } else if (months <= 12){
        return months + "月前";
    }else return timeString;
  }


export { FormatDateTime, GetTimePeriod}