
////////////////////////////////////////////工具对象/////////////////////////////////////////////////////////////////
////////////////////////////////////////////工具对象/////////////////////////////////////////////////////////////////
////////////////////////////////////////////工具对象/////////////////////////////////////////////////////////////////
function Utils() {}

Utils.prototype.getSplitSongCoverName = function (str,endIndex) {
    if (endIndex >= str.length){
        return str
    }
    if (str.length <7){
        return str
    }
    let result = str.slice(0,endIndex) + '...';
    return result
};