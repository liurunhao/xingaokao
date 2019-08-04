/*会员登录*/
 
  

  function user_submit_but()
    {
      flag=1;
      if($('#user_name').val()=='')
      {
         alert('请输入用户名！');flag=0;
      }else if($('#password').val()=='')
      {
         alert('请输入密码！');flag=0;
      }
      if(flag==1)
      {
          $.ajax({    
                url:'/login/sign_in_ajax',
                data:{  
                        user_name : $('#user_name').val() ,  
                        password : $('#password').val()
                },    
                type:'post',     
                dataType:'html', 
                async:false ,  
                beforeSend:function(data) {    
                     $('#user_submit_but').attr('disabled','disabled');
                     $('#user_submit_but').html('会员登录中，请稍后...');
                 },    
                success:function(data) { 
                    if(data=='ok')
                    {window.location.reload();}
                    else
                      {
                        alert(data);
                        $('#user_submit_but').removeAttr('disabled');
                        $('#user_submit_but').html('登录');
                      }
                 },    
                 error : function() {    
                      alert("error!");    
                 }    
               }); 
        }  
  }

 
/*会员注册*/
 
		/*获取验证码*/
		function get_invalid_code(){
			var user_name = $("#r_user_name").val();
			if(!isNaN(user_name) && user_name.length == 11){
				var query_url = '/login/get_verifi_code';
				var send = {'phone':user_name};
                $.ajax({    
	              url:query_url,
	              data:send,    
	              type:'post',     
	              dataType:'html', 
	              async:false ,  
	              beforeSend:function(data) {   
	                   $("#get_invalid_code").attr('disabled','disabled');
	                   $("#get_invalid_code").html('短信发送中');  
	               },    
	              success:function(data) { 
	                  if(data=='ok')
	                  {
	                  	alert('短信验证码已发送，请查收！');
	                  	$("#get_invalid_code").html('接收(<span></span>)'); 
	                  	$("#get_invalid_code").css('cursor','text');
	                  	$("#get_invalid_code").attr('disabled','disabled');
	                  	down_second(120); 

	                  }
	                  else
	                  {
	                  	alert(data);
	                  	$("#get_invalid_code").html('获取验证码'); 
                        $("#get_invalid_code").removeAttr('disabled');
	                   }
	               },    
	               error : function() {    
	                    alert("error!");    
	               }    
	             });
 		
			}else{
				alert('请输入正确的手机号');
			}
		} 
		
		/*提交*/
		function r_user_submit_but(){
			var user_name = $("#r_user_name").val();
			var invalid_code = $("#r_invalid_code").val();
			var password = $("#r_password").val();
			if('' == user_name){
				alert('请输入正确的手机号');
				return false;
			}
			else if('' == invalid_code){
				alert('请输入验证码');
				return false;
			}
			else if('' == password){
				alert('请输入新的密码');
				return false;
			}else if(password.length < 6){
				alert('请输入6位以上的密码');
				return false;			
			}
			else
			{
                $.ajax({    
	              url:'/login/register_ajax',
	              data:{  
                      user_name : $('#r_user_name').val() ,  
                      password : $('#r_password').val(),
                      invalid_code :  $("#r_invalid_code").val()
                      },    
	              type:'post',     
	              dataType:'json', 
	              async:false , 
	              beforeSend:function(data) {    
                   $('#r_user_submit_but').attr('disabled','disabled');
                   $('#r_user_submit_but').html('会员注册中，请稍后...');
                   }, 
	              success:function(data) { 
	                  if(data.msg=='ok' )
	                  {
	                  	alert('注册成功！');
                        window.location.reload();
	                  }
	                  else
	                  {
                         alert(data.msg);
                         $('#r_user_submit_but').removeAttr('disabled');
                         $('#r_user_submit_but').html('注册');
	                  }
	               },    
	               error : function() {    
	                    alert("error!");    
	               }    
	             });
			}
 
		}
	 
 
function down_second(intDiff){
	$("#get_invalid_code span").html(intDiff);
	intDiff--;
	if(intDiff>0) {
		var t=setTimeout("down_second("+intDiff+")",1000);
	} else{
		clearTimeout(t);
		$("#get_invalid_code").html('获取验证码'); 
        $("#get_invalid_code").removeAttr('disabled');
	}   
} 