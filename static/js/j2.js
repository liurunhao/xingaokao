/*��Ա��¼*/
 
  

  function user_submit_but()
    {
      flag=1;
      if($('#user_name').val()=='')
      {
         alert('�������û�����');flag=0;
      }else if($('#password').val()=='')
      {
         alert('���������룡');flag=0;
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
                     $('#user_submit_but').html('��Ա��¼�У����Ժ�...');
                 },    
                success:function(data) { 
                    if(data=='ok')
                    {window.location.reload();}
                    else
                      {
                        alert(data);
                        $('#user_submit_but').removeAttr('disabled');
                        $('#user_submit_but').html('��¼');
                      }
                 },    
                 error : function() {    
                      alert("error!");    
                 }    
               }); 
        }  
  }

 
/*��Աע��*/
 
		/*��ȡ��֤��*/
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
	                   $("#get_invalid_code").html('���ŷ�����');  
	               },    
	              success:function(data) { 
	                  if(data=='ok')
	                  {
	                  	alert('������֤���ѷ��ͣ�����գ�');
	                  	$("#get_invalid_code").html('����(<span></span>)'); 
	                  	$("#get_invalid_code").css('cursor','text');
	                  	$("#get_invalid_code").attr('disabled','disabled');
	                  	down_second(120); 

	                  }
	                  else
	                  {
	                  	alert(data);
	                  	$("#get_invalid_code").html('��ȡ��֤��'); 
                        $("#get_invalid_code").removeAttr('disabled');
	                   }
	               },    
	               error : function() {    
	                    alert("error!");    
	               }    
	             });
 		
			}else{
				alert('��������ȷ���ֻ���');
			}
		} 
		
		/*�ύ*/
		function r_user_submit_but(){
			var user_name = $("#r_user_name").val();
			var invalid_code = $("#r_invalid_code").val();
			var password = $("#r_password").val();
			if('' == user_name){
				alert('��������ȷ���ֻ���');
				return false;
			}
			else if('' == invalid_code){
				alert('��������֤��');
				return false;
			}
			else if('' == password){
				alert('�������µ�����');
				return false;
			}else if(password.length < 6){
				alert('������6λ���ϵ�����');
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
                   $('#r_user_submit_but').html('��Աע���У����Ժ�...');
                   }, 
	              success:function(data) { 
	                  if(data.msg=='ok' )
	                  {
	                  	alert('ע��ɹ���');
                        window.location.reload();
	                  }
	                  else
	                  {
                         alert(data.msg);
                         $('#r_user_submit_but').removeAttr('disabled');
                         $('#r_user_submit_but').html('ע��');
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
		$("#get_invalid_code").html('��ȡ��֤��'); 
        $("#get_invalid_code").removeAttr('disabled');
	}   
} 