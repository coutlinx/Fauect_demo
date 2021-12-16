$('.navTrigger').click(function () {
  $(this).toggleClass('active')
  console.log('Clicked menu')
  $('#mainListDiv').toggleClass('show_list')
  $('#mainListDiv').fadeIn()
})

function Value(e) {
  let value = e.innerHTML.split(' ETH')[0]
  $('#value').val(value)
  but = $('#Btn')
    $('#Btn').addClass("btn-info").removeClass("btn-outline-info")

  but.text(e.innerHTML)
  // `<button type="button" class="btn btn-lg btn-primary dropdown-toggle" disabled>${e.innerHTML}</button>`
}

$('#Send').click(() => {
  if ($('#value').val() != 0) {
      if ($('#Address').val() == 0){
        confirm('请确认您的账户')
        return
      }
    Address = $('#Address').val()
      Value = $('#value').val()
      value = Value*10**18
      $.ajax({
          url:"http://localhost:1234/linx/transfer",
          type:"post",
          dataType:"json",
          data:{address:Address,amount:value},
          success:function (data) {
              console.log(data)
          },
          error:function (data){
              console.log(data)
          }
      })
    console.log($('#value').val(), Address)
  } else {
    confirm('请确认您的ETH数目')
    return
  }
})
$('#balance').click(()=>{
  $.ajax({
      url:"http://localhost:1234/linx/cAmount",
      type:"post",
      dataType:"json",
      success:function (data){
        console.log(data)
        $("#exampleModalLongTitle").html("合约余额")
        $(".modal-body").html(`<div class="jumbotron">
                      <h1 class="display-3">合约余额:</h1>
                      <hr class="my-5">
                      <p>${data.data}</p>
                      <br/>
                      <p class="lead">
                        <a class="btn btn-primary btn-lg" href="https://github.com/coutlinx/Fauect_demo" role="button">查看合约代码</a>
                      </p>
                    </div>`)
        $("#exampleModalLong").modal("show")
      },
      error:function (data){
        console.log(data)
      }
  })

})
$("#query").click(()=>{
    $("#exampleModalLongTitle").html("查询用户领取数量")
      $(".modal-body").html(`<div class="jumbotron">
  <h1 class="display-3"><div class="input-group mb-3">
  <input type="text" class="form-control" placeholder="输入要查询的账户" id= "Qadd" aria-label="Recipient's username" aria-describedby="basic-addon2" style="Height:60px">
  <div class="input-group-append">
    <button class="btn btn-info" type="button" style="width:"60px"" onclick = "Query()">Button</button>
  </div>
</div></h1>
  <hr class="my-5">
  <p id="VALUE">0</p>
  <br/>
  <p class="lead">
    <a class="btn btn-primary btn-lg" href="https://github.com/coutlinx/Fauect_demo" role="button">查看合约代码</a>
  </p>
</div>
`)
  $("#exampleModalLong").modal("show")
})

function  Query(){
    address = $("#Qadd").val()
    console.log((address))
    $.ajax({
        url:"http://localhost:1234/linx/getLimit",
        type:"post",
        dataType:"json",
        data:{address:address},
        success:function (data) {
            console.log(data.data)
            $("#VALUE").html(data.data)
        },
        error:function (data){
            console.log(data)
        }
    })
}