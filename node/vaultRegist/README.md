# 加密模式

使用"xrsa"库里的RSA加密函数创建公钥私钥保存为public.pem和private.pem两个文件以便后续的使用。在发送过程中默认Client已知Node端的公钥，Client发送的消息还需要附带自己的公钥。Client端需要发送的消息格式为：

{

  "Email": "" ,

  "Phone": "" ,

  "DID": "" ,

  "Pubkey": ""

}

4个参数均为string类型。

其中DID是经过双重加密的document文档的密文。首先用Client的私钥进行加密获得密文encry1,然后用Node端的公钥加密获得DID。

以下是Client端发送的消息实例：

{

  "Email": "123456@gmail.com" ,

  "Phone": "123456789" ,

  "DID": "Q3h3e4mhyAjOgRJfRe7z0cti53U6-g9ADjbK4oFaCAYE-WALV2CTwT7-xqxi2P8d5O94q9NlNgsmjJm5XHuFE-26D8z8H-TilU1n_YANejicmJ_jH5vVAKIA8V98uJLXQOmEpnjw3Anl-qfCPokBZahFZ9AUNm5XrQeO6GR0C7s3u8BD3T-azju1Rk5NeseCHCusBHfW8i1VsWm2coRibzREGslyiRjiV-QUviOwLxWvo3E43QJllSJIXLCmsAYEZrqi8Q2ebOuEUJpuslsHyxcdJkXceD4gr-Mu-KgKEzKp76Umg5_tBifyLgKw2ZF1AT_rI0OLVtG_7_JmNcNemBOL2uhneQSX_A_LseC-a06DhI9RVhbYjooPNofDcmqtInfGj6gdyU4mCyrSQggWhrmsuBtIBIEt3_wF4ipyIE7bP282Qlcfeu368WSxu4u_w9gK1I7ol31R_Mp_ISV_Ni_XcSMdWQhG14ic3hsA-axVUJmxbKR3C4eZJMgnyXd1JleGyihhGu8VOdaXW4WVPNkPkv-vJZAHBr_I5z7fPDIYJhXNd5S2GU79GrNZ2JxynpYgxIfbedthhV2g_WWv11g0lp6laHx1oLbCQpGqvA-lLBG3yrMAZzOW7qp-6rgl49wmTtGDGYNNQtJgjtnDrT1dL6HsgC-no2SSvyJ3tNVjK2_9vCMqZYvNCxV40gpxZml8CbLipnMCJJYD7lcZwDOZZX6FWze9dnOli-DAkRnkCIA1RdocA2VAv8k6qRGmMvLUT9Dg6aRPt8ZtMekUuTYEWcbbgR_9Mm7jjGp8gmQ8mJV_Q_XgRxySPb64z59qDN9kOJzhjYaNCp_KvmolILJWGPYYAK8PuE7X2MLNB5b5h5qJy7ntf5JKCbp79zT8q13yJ61u-M1bODiszsRHQH8xX1042yx30mX1RPIQvUBN_pcmYlOfKir3_BVm4XlU5c2Kf4k8L74Q2HUf7I-Cmi7WLaLPMoXiTES2JN66puRw7vk5CM5y1fQvrL6eteAXiUCarXQ7vPKBedRSj0E4jh5nqO1YlgwXHqSCtnYGuIo8T1WM8hplfvwx-tUwN-4oruwomJSF0bpu67ll3Hof0Dxbtfd53yvX5Q4hhQL_eelLOl6qvK7RQXMyV9UdxOe7eFE7tlM8I8d-r07woPZCyKwAd-FS6WRTi1jyvYnHEz4sjnLHPbl6eGU9mfkZ5mVWZ8jmOy0r_7-7AbsRIyl9mfbA-oSc4UhxcBsYY_2thjMaV1f9DNBmfNVcGSfICrIBNzO0xQxIrBH81OBlABTqwJcNHIccUNAKLgaRJG32VmtNnLNFhWFlh-RcVD8S_Ez6qrVfAmnpp3D4qPVEfbfwSoc9AfakAE_1jv-p0nlZ-tssZ6vALVQOwgXm7ul2-HVqswm1zB3tTd9AnJEc4M5cIqjyAN365-aPCepy31m4AM7Z8oiCrdoGOJMv64mxa1yl_4agPt3JhZfxOWcgErYcEDYu6ZkjAUXHyfIGpMuyMHpuItFjW22faGbmVMD0WtMpDnMHZy96MiF3mVMNFxQ5XVQhG9pC_HAwH_D6jpQGLzMfM4ffPl9SNcz4K7RCxOE5VKqeq1PwWX5jkUm2OmrT6tEeJefRynKRd0aQW9lDlpbpxzK5vfu4NlGAOJB8m_CWnrNgflQtwo2fwVOzOuohFCGT1hd-JjkgJNtFj2xXl_tv5wX3V7SoX-xCQmh5p8nC3UNLLs3JwlDdGTSvLV63tJHDJM23C6HUpb5E9MG0kqh6m7v-FzzUekluGnhjxWYfgcXGBWc1c-CSzESQ2VjXsEu__ss95yVUod3HKzoO8i4wdXBPiXc9Qu4d_JMdMTz8dtvZjlEnavfyOYKeP0XYGwmUc5LBuGD-CfWctKAZCaFXFIikG5xQ_zn9PxltdZWgA3HTYZyKGSJl_Pip0DzMIFBGbVVjlTl9j3_XxA16Xm2NKTqHrWkrw2MUe-bSkfbyxMEOuGcXeSDDdSCaAcAl8jrcbz3auS5V8aCJNxwKZgMI63vsj5vPqMQyNQEAMa66PfH_0CJf265UnMrmLT71fSeCZ3bdzvOPuL9dfJHTGIsN7Vd-jJBCqBlyuI_eaCWkywYDZu-TAzqqshIGVx8MzWiyWuXzIOiN9O9qj2mRGNA9_vbq0ezEwrUGQU5YndI86GJTPGKjQYu-xCFpbi9HXQ2Ht-32mpPCbB1C2gshSXR9ehcZ1k_fJuRBFkIHNgNagPdhAcSl9ZKjUhDaByJnKcGPyjBMWaKoyPDU5TR5_troEX03Bu_pifv7yuNIWr_kn9UfmT4HSf5TbbVNvxywqAUlk5FJGYLBOPRLqR3E4jOqyAwdmKttMWMp5bu7PHkxdwYGpTB1h8GrmjCFK6tt-EidMQ_bqdakuE--5YZZUYKXC_Do8MGjjaDhajUGDkr-G8Xqre699h4tY3A3OLQhDGZ01A34OEDw3hYqSoHNurAhyUatG2ZEtO2ueS4PZQhcTyhwQ8_1jNwGDhz_6RdCQGOF7TQFi7OfGPeY0h-WMWDknCdSo50jBPUm1nQ5q9ceXjQgwuEzRX4wJXZgFmE41BRrWcpTTqt0UoDZxGteMuB6fgiTJs-t4uVKkabL5H65rrNr6Pu4Qeeyq0bpkKDBYkWITVEvJ76IVyJYX4xbgzVTPHL_GK9_wISx0wKF8VdHxDh-26TeogE3hQSEAmNNbtReBNRN35EJNyZMhehlBnxXigSJjLf6BdQ_H8VD_O4jKgPxt5dJPEsygYpK2zPnIYeamT-5rlVq8xqcwEX9RrDGrt5mbF-uqU12M9AkPBKIlUCYCTmMFqySlM1vgNUR3FKYmfiQwJXY0pebg6FpWxNqCVGlyfYNbJN3jHtFp-tk3J91VVoVLwLTteU52j4pM1W35bgfs2xcaqJxTpM3sEZ-C5STkxCxmKyLCafTF8UI0O46yM73mSfC0jOCPqIE4zyYINrjKmz7s5w5QbvrcZELueVozKFR5LK0MnsPGonkD6s2CalZC2-RcOmNrvJPCwB0wK8vgXXm7SIjT4rF5QcNJQ0e8G2BHoqujvunx1emEErcKBbYYA9ITE1T0eFmj0enKECJhAPUrc4MUuqFoVULuj0lhENJ5UGPc3KY81iShrMtzwJ41Z-D7AVDhaLi3iaPpYJLALk7aOVQ-N9RhaxYBQi0PTuCc57IocrOapNN91pusHUFId2g87vIsuJVfZ4pdmw3WNnGUTPBTdbfICn3dMwUJLFv-fVdRYLNRDlSeXmzV67n9GPFNs7Tu0bLBE52YDe-kA9gK3MyKR-5kXcnw0swcJzVDuUy8rqqysD5gpAYBxDUOLQ065qRbLwVrk42sVQPYiHjAI6iaaFI0W_ajxqS-uyelocS00hcb4H9EHakpo8YXvYIR5svhiboNA" ,

  "Pubkey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoTdbcYcLggJjzpQ/0b4A\n67RDQqvZTvAAfE9K8vilwOs0AkMX1RWm3GqIGCrexc9clW0a9uGw9XwFiB7HNfcB\nGhIMIa+5Q5nczbPTFESl4ssLkM9TPBCSJ4Pt9fEEF5TwBgFirBqQ6CMd0U2jg24y\n8YFuTdFMfiSAixReahlIu7lyaMbR8oG3R6ecTbCC6P/5GM9pIeQFWHN3wwf1YQqf\naleWT/TcUJf5NL0o3E49guoS0c3ZMBNkp2A2vO2BGXb6w2HhdNfja0CdiH4Tx+rS\nLby9IAIGvxCu21ooazXocv2XcM8n8YG5Y2B4zdj2S7RNda0QVAecWcejh0tfP60D\ntwIDAQAB\n-----END PUBLIC KEY-----\n"

}

# 不加密模式

该模式下直接将DID document发送即可。例如：

{

 "@context": [

  "https://www.w3.org/ns/did/v1"

 ],

 "id": "did:nuts:04cf1e20378a4e38ab1b401a5018c9ff",

 "controller": [

  "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff",

  "did:nuts:f03a00f1-9615-4060-bd00-bd282e150c46"

 ],

 "verificationMethod": [

  {

   "id": "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff",

   "type": "JsonWebKey2020",

   "controller": "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff",

   "publicKeyJwk": {

​    "kty": "EC",

​    "crv": "P-256",

​    "x": "SVqB4JcUD6lsfvqMr-OKUNUphdNn64Eay60978ZlL74",

​    "y": "lf0u0pMj4lGAzZix5u4Cm5CMQIgMNpkwy163wtKYVKI"

   }

  },

  {

   "id": "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-2",

   "type": "JsonWebKey2020",

   "controller": "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff",

   "publicKeyJwk": {

​    "kty": "EC",

​    "crv": "P-256",

​    "x": "SVqB4JcUD6lsfvqMr-OKUNUphdNn64Eay60978ZlL74",

​    "y": "lf0u0pMj4lGAzZix5u4Cm5CMQIgMNpkwy163wtKYVKI"

   }

  },

  {

   "id": "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff#added-assertion-method-1",

   "controller": "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff",

   "publicKeyBase58": "GGRj8PAR5tRgD5xqAhPna1bLa3UoYuxNEEhRmcYCPBm5",

   "type": "Ed25519VerificationKey2018"

  }

 ],

 "authentication": [

  "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff",

  "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff"

 ],

 "assertionMethod": [

  "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff"

 ],

 "service": [

  {

   "id": "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff#service-1",

   "type": "nuts:bolt:eoverdracht",

   "serviceEndpoint": "did:nuts:<vendor>#service-76"

  },

  {

   "id": "did:nuts:04cf1e20-378a-4e38-ab1b-401a5018c9ff#service-2",

   "type": "nuts:core:consent",

   "serviceEndpoint": "did:nuts:123456"

  }

 ]

}