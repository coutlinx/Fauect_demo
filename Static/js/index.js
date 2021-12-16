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
              if (data.code == 1){
              console.log(data.message)
                  if (data.message.code == -32000){
                      $("#exampleModalLongTitle").html("出了点问题")
                      $(".modal-body").html(`<h1 class="text-danger  text-center">每个账户每天只能领取3ETH哦</h1>`)
                      $("#exampleModalLong").modal("show")
                      return
                  }else if(data.message.code == undefined){
                      $("#exampleModalLongTitle").html("出了点问题")
                      $(".modal-body").html(`<h1 class="text-danger  text-center">${data.message}</h1>`)
                      $("#exampleModalLong").modal("show")
                      return
                  }
              }else {
                  console.log(data)
              }
          },
          error:function (data){
              console.log(data)
              $(".alert").alert()
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
          if (data.code ==1){
              if (data.message.URL != undefined)
              $("#exampleModalLongTitle").html("出了点问题")
              $(".modal-body").html(`<h1 class="text-danger  text-center">私有链没有链接上</h1>`)
              $("#exampleModalLong").modal("show")
              return
          }
        $("#exampleModalLongTitle").html("合约余额")
        $(".modal-body").html(`<div class="jumbotron">
                      <h1 class="display-3">合约余额:</h1>
                      <hr class="my-5">
                      <p class="text-info  text-center" style="font-size: 32px">${data.data/10**18}  ETH</p>
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
            console.log(data)
            if (data.code ==1){
                if (data.message.URL != undefined)
                    $("#exampleModalLongTitle").html("出了点问题")
                $(".modal-body").html(`<h1 class="text-danger  text-center">私有链没有链接上</h1>`)
                $("#exampleModalLong").modal("show")
                return
            }
            $("#VALUE").html(data.data)
        },
        error:function (data){
            console.log(data)
            $(".alert").alert()
        }
    })
}
