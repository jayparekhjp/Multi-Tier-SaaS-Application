/*$(document).ready(function(){
     var userid = $("#userid").get(0).getAttribute('data-info');
     var restraunt_id = $('#count').get(0).getAttribute('res_id');
     $.ajax({
       type: 'GET',
       url: 'http://demo8655652.mockable.io/count',
       data: { 
        userid: userid, 
        restraunt_id : restraunt_id 
       },
       dataType: "json",
 
       crossDomain: true,
       success: function (msg) {
         console.log(msg);
         msg.forEach(function(value) {
           if($('p[item_id='+value.item_id+']').get(0) !== undefined){
             $('p[item_id='+value.item_id+']').get(0).innerHTML = 'x '+value.count;
           }
         }); 
       },
       error: function (request, status, error) {
 
          alert(error);
       }
    });
 });
 */
 $(".add-cart").on("click",function(){
   var item_id = this.value;
   var restraunt_id = this.getAttribute("res_id");
   var item_name = this.getAttribute("item_name");
   var item_price = parseFloat(this.getAttribute("item_price"));
   var restraunt_name = $("#restraunt_name").text();
   var userid = $("#userid").get(0).getAttribute('data-info');

   var res_name = restraunt_name.trim();
   /*console.log("res_id:"+this.getAttribute("res_id"))
   console.log("item_id:"+this.value)
   console.log("THIS WORKS")*/
   $.ajax({
     type: 'POST',
     url: 'http://34.219.121.214:3000/api/cart/itemSave',
     data: JSON.stringify({ 
      rid: restraunt_id, 
      iid: item_id,
      iname : item_name,
      price : item_price,
      res : res_name,
      id : userid 
     }),
     dataType: "json",
 
     crossDomain: true,
     success: function (msg) {
        console.log(msg)
     },
     error: function (request, status, error) {
 
        alert(error);
     }
  });
 });
 
 $("#checkout").on("click",function(){
  window.location.href = "/viewcart"
 });
 