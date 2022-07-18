import UIKit
import Goprojlib

class ViewController: UIViewController {
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        DispatchQueue.main.asyncAfter(deadline: .now()+1) {
            
            /* Calling the test function, off the main thread */
            DispatchQueue.global(qos: DispatchQoS.QoSClass.userInitiated).async {
//                Goprojlib.GoprojlibRunChannels()
//                Goprojlib.GoprojlibRunChannelsMinimizing()
                Goprojlib.GoprojlibRunChannelConciseExample(false)
            }

        }
    }
    
}
